import React, { useState, useEffect } from 'react'
// import { Form, Upload, Button, InputNumber } from 'antd'
// import { UploadOutlined } from '@ant-design/icons'
import { Input, Row, Col } from 'antd'

import axios from 'axios';
import { config } from '../config';

const validImageTypes = ['image/gif', 'image/jpeg', 'image/png'];

export default function FormTest(props) {

  const [ id, setId ] = useState('')
  const [ file, setFile ] = useState({})
  const [ validated, setValidated ] = useState(false)

  useEffect(() => {
    if (file === undefined || id === undefined) {
      setValidated(false)
    } else setValidated(validImageTypes.includes(file.type) && id !== '')
  }, [id, file])

  const handleSubmit = async (event) => {
    event.preventDefault();
    if (!validated) {
      event.stopPropagation()
      console.log("Need to fill in!")
      return
    }

    // create data
    let data = new FormData()
    await data.append('file', file)
    await data.append('id', id)
    const requestConfig = {     
      headers: { 'content-type': 'multipart/form-data' }
    }

    // Display the key/value pairs
    // for (var pair of data.entries()) {
    //   console.log(pair); 
    // }

    // upload the file
    await axios.post(
      `${config.IMAGE_BACKEND_SERVER}`,
      data, 
      // {'content-type': `multipart/form-data; boundary=${data._boundary}`}),
      requestConfig,
    ).then(res => {
      console.log(res)
    }).catch(error => {
      console.log("Err: " + error)
    })
  }

  return (
    <>
      <form onSubmit={handleSubmit}>
        <label htmlFor="id">ID:</label><br/>
        <Input type="number" id="id" name="id" style={{marginBottom: "5px"}}
          onChange={(e) => {
            setId(e.target.value)
          }} value={id}/> <br/>
        <label htmlFor="picture">Picture:</label>
        <Input type="file" id="picture" name="file" style={{marginBottom: "15px", border: "none"}}
          onChange={(e) => {
            setFile(e.target.files[0])
          }}/>
        <Row>
          <Col span={8}>
            <Input type="submit" value="Submit" style={{marginBottom: "15px", borderColor: "black"}}/>
          </Col>
        </Row>
      </form>

      {/* <Form
        {...layout}
        layout="horizontal">
          
          <Form.Item label="Product ID">
            <InputNumber onChange={(value) => {
              setId(value)
              setValidated(validImageTypes.includes(file.type) && id !== '')
            }} value={id}/>
          </Form.Item>

          <Form.Item
            label="File"
            valuePropName="fileList"
            getValueFromEvent={normFile}
            extra="Image to associate with the product">
            <Upload {...uploadProps}>
              <Button icon={<UploadOutlined />}>Choose file</Button>
            </Upload>
          </Form.Item>

          <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 5 }}>
            <Button onClick={handleSubmit}>Submit</Button>
        </Form.Item>
      </Form> */}
    </>
  )
}
