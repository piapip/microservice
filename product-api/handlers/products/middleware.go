package handlers

import (
	"context"
	"net/http"

	"github.com/piapip/microservice/product-api/data"
)

// MiddlewareProductValidation will extract the product struct from req, fill the empty struct with data then put it into req's Context() with Value of newProduct and key as KeyProduct{}
// So next time if you want the data from req, you call req.Context() *since we tuck the data in context
// Then call the parameter .Value() since we assign it with WithValue()
// And specify the key we assign the Value to, in this case is KeyProduct{}. This way, we can bind a lot of value into req.Context.Value().
// Example: req.Context().Value(KeyProduct{}).(data.Product{}) <- the .(data.Product{}) is to convert the extracted data to Product struct.
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		// you can assign newProducts to data.Product{} (without &) but you'll also have to update other codes related to it.
		// For example: models.AddProduct(newProduct)
		//  will become models.AddProduct(newProduct).
		// Either way is fine. Value the readability.
		newProduct := &data.Product{}

		// extracting data from req to JSON
		// if you change newProduct(equivalent to &data.Product{}) to *newProduct(equivalent to data.Product{})
		// it will tell you that you fuck up the type.
		err := data.FromJSON(newProduct, req.Body)
		if err != nil {
			p.logger.Error("Unable to deserialize products", "error", err)

			// http.Error(res, "Something's wrong with the server. Unable to convert from JSON to Product struct", http.StatusBadRequest)
			res.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, res)
			return
		}

		// validate the product
		// it's errs, it will return error if you use "err" instead
		errs := p.validator.Validate(newProduct)
		if len(errs) != 0 {
			p.logger.Error("Unable to validate product", "error", errs)

			// return the validation messages as an array
			// http.Error(res, fmt.Sprintf("Error validating product: %s", errs), http.StatusUnprocessableEntity)
			res.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, res)
			return
		}

		// add the product to the context
		// As said above, you can make life of those who's writing REST easier by doing a conversion on this side.
		// You can put it as ctx := context.WithValue(req.Context(), KeyProduct{}, newProduct) but
		// for example, the post.go, it must be newProduct := req.Context().Value(KeyProduct{}).(*data.Product) instead.
		// Since our newProduct is &...
		ctx := context.WithValue(req.Context(), KeyProduct{}, *newProduct)
		req = req.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(res, req)
	})
}
