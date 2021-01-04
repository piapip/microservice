import HomePage from './containers/HomePage'
import Admin from './containers/Admin'
import 'antd/dist/antd.css';

import { BrowserRouter, Route } from 'react-router-dom';

function App() {
  return (
    <BrowserRouter>
      <Route exact path='/' 
        render={(props) => {
          return <HomePage />
        }}/>

      <Route exact path='/admin'
        render={(props) => {
          return <Admin />
        }}/>

    </BrowserRouter>
  );
}

export default App;
