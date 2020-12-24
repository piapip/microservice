import React from 'react'
import { Table } from 'antd'

export default function ItemList({ items }) {

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      align: 'center',
    },
    {
      title: 'SKU',
      dataIndex: 'sku',
      align: 'center',
    },
  ]

  return (
    <Table 
      dataSource={items} columns={columns} rowKey="id"
      pagination={{ pageSize: 10, position: ["none", "bottomCenter"] }} />
  )
}
