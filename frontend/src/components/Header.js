import React, { useState } from 'react'
import { Link } from 'react-router-dom'

import { Menu } from 'antd'

export default function Header() {

  const [ currentTab, setCurrentTab ] = useState('')
  const { SubMenu } = Menu;

  const handleClick = (e) => {  
    setCurrentTab(e.key)
  }

  return (
    
    <Menu onClick={handleClick} selectedKeys={[currentTab]} mode="horizontal">
      {/* <Menu.Item>Coffee Shop</Menu.Item> */}
      <SubMenu title='Coffe Shop'>
        <Menu.Item key="setting:1"><Link to='/'>Home</Link></Menu.Item>
        <Menu.Item key="setting:2"><Link to='/admin'>Admin</Link></Menu.Item>        
      </SubMenu>
      <Menu.Item key='Home'><Link to='/'>Home</Link></Menu.Item>
      <Menu.Item key='Link'><Link to='/admin'>Admin</Link></Menu.Item>  
    </Menu> 
  )
}
