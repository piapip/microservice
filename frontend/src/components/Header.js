import React, { useState } from 'react'

import { Menu } from 'antd'

export default function Header() {

  const [ currentTab, setCurrentTab ] = useState('')
  const { SubMenu } = Menu;

  const handleClick = (e) => {
    console.log(`click: ${e.key}`)
    setCurrentTab(e.key)
  }

  return (
    
    <Menu onClick={handleClick} selectedKeys={[currentTab]} mode="horizontal">
      {/* <Menu.Item>Coffee Shop</Menu.Item> */}
      <SubMenu title='Coffe Shop'>
        <Menu.Item key="setting:1">Home</Menu.Item>
        <Menu.Item key="setting:2">Admin</Menu.Item>        
      </SubMenu>
      <Menu.Item key='Home'>Home</Menu.Item>
      <Menu.Item key='Link'>Link</Menu.Item>  
    </Menu> 
  )
}
