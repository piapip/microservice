import React, { useState } from 'react'

import { Menu } from 'antd'

export default function Header() {

  const [ currentTab, setCurrentTab ] = useState('')

  const handleClick = (e) => {
    console.log(`click: ${e.key}`)
    setCurrentTab(e.key)
  }

  return (  
    <Menu onClick={handleClick} selectedKeys={[currentTab]} mode="horizontal">
      <Menu.Item>Coffee Shop</Menu.Item>
      <Menu.Item key='Home'>Home</Menu.Item>
      <Menu.Item key='Link'>Link</Menu.Item>
    </Menu> 
  )
}
