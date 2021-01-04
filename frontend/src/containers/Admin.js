import React from 'react'
import { Row, Col } from 'antd'

import Header from '../components/Header'
import FormTest from '../components/FormTest'

export default function Admin() {

  return (
    <>
      <Header />
      <h1 style={{marginBottom: "30px", marginTop: "40px", textAlign:'center'}}>Admin</h1>

      <Row>
        <Col span={8} offset={8}>
          <FormTest />
        </Col>
      </Row>
    </>
  )
}
