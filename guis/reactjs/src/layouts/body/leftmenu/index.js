import React from 'react';
import {
  CDBSidebar,
  CDBSidebarContent,
  CDBSidebarFooter,
  CDBSidebarHeader,
  CDBSidebarMenu,
  CDBSidebarMenuItem,
} from 'cdbreact';
import { NavLink } from 'react-router-dom';

export class LeftMenu extends React.Component {
  render() {
    return (
        <div id="leftmenu" style={{ height: '100vh', overflow: 'scroll initial' }}>
            <CDBSidebar textColor="#fff" backgroundColor="#333">
                <CDBSidebarHeader prefix={<i className="fa fa-bars fa-large"></i>}>
                    <a href="/" className="text-decoration-none" style={{ color: 'inherit' }}>
                        Menu
                    </a>
                </CDBSidebarHeader>

                <CDBSidebarContent className="sidebar-content">
                    <CDBSidebarMenu>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Home</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Dashboard</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Profiles</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Components</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Responses</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Requests</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Instances</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Databases</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Queries</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="columns">Composes</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/" className="activeClicked">
                            <CDBSidebarMenuItem icon="columns">Grammars</CDBSidebarMenuItem>
                        </NavLink>
                    </CDBSidebarMenu>
                </CDBSidebarContent>

                <CDBSidebarFooter style={{ textAlign: 'center' }}>
                  <div
                    style={{
                      padding: '20px 5px',
                    }}
                  >
                    Sidebar Footer
                  </div>
                </CDBSidebarFooter>
        </CDBSidebar>
    </div>
    );
  }
}
