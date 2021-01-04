import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { Row, Col } from 'antd'

import Header from '../components/Header'
import ItemList from '../components/ItemList'
import SearchBar from '../components/SearchBar'
import { config } from '../config'

export default function HomePage() {

  const [ items, setItems ] = useState([])

  const [ searchTarget, setSearchTarget ] = useState('')

  // const getAllItems = () => {
  //   axios.get(`${config.BACKEND_SERVER}/products`)
  //     .then(async (response) => {
  //       await setItems(response.data)
  //       console.log(response.data)
  //     })
  //     .catch(err => {
  //       console.log(err)
  //     }) 
  // }

  useEffect(() => {
    const getAllItems = async () => {
      await axios.get(`${config.BACKEND_SERVER}/products`)
        .then(async (response) => {
          await setItems(response.data)
        })
        .catch(err => {
          console.log(err)
        }) 
    }

    getAllItems()
    
  }, [searchTarget]);

  const showItems = items.filter(item => {
    return item.name.includes(searchTarget)
  })

  return (
    <>
      
      <Header />
    
      <h1 style={{marginBottom: "10px", marginTop: "40px", textAlign:'center'}}>Menu</h1>
      <Row style={{marginBottom: "30px"}}>
        <Col span={8} offset={8}>
          <SearchBar 
            setSearchTarget={setSearchTarget}/>
        </Col>
      </Row>
      
      <Row>
        <Col span={12} offset={6}>
          <ItemList items={showItems} />
        </Col>
      </Row>

      {items ? <span>{items.name}</span> : null}
    </>
  )
}
