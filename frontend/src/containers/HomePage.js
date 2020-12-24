import React, { useState, useEffect } from 'react'

import Header from '../components/Header'
import ItemList from '../components/ItemList'
import axios from 'axios'
import config from '../config'

export default function HomePage() {

  const [ items, setItems ] = useState([
    {
      test1: "test1.1",
      test2: "test2.1",
    },
    {
      test1: "test1.2",
      test2: "test2.2",
    },
  ])

  const getAllItems = () => {
    axios.get(`${config.BACKEND_SERVER}/products`)
      .then((response) => {
        console.log(response)
      })
      .catch(err => {
        console.log(err)
      }) 
  }

  useEffect(() => {
    getAllItems()
  }, []);

  return (
    <>
      <Header />
      <h1 style={{marginBottom: "40px", marginTop: "40px", textAlign:'center'}}>Menu</h1>
      <ItemList items={items} />
    </>
  )
}
