import React from 'react'
import { Input } from 'antd'

// props needs to have 2 things: searchTarget and setSearchTarget
export default function SearchBar(props) {

  const updateFilter = (e) => {
    props.setSearchTarget(e.target.value)
  }

  return (
    <Input placeholder="This is a filterer." onChange={updateFilter}/>
  )
}
