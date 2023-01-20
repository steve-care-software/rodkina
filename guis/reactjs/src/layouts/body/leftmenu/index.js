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

import './index.css';

export class LeftMenu extends React.Component {
  render() {
    return (
        <div id="leftmenu" style={{ height: '100vh', overflow: 'scroll initial' }}>
            <CDBSidebar textColor="#fff" backgroundColor="#333">
                <CDBSidebarHeader prefix={<i className="fa fa-bars fa-large"></i>}>
                    <a href="/" className="text-decoration-none" style={{ color: 'inherit' }}>
                        WebX Network
                    </a>
                </CDBSidebarHeader>

                <CDBSidebarContent className="sidebar-content">
                    <CDBSidebarMenu>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-house" iconSize="lg">Home</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-bar-chart" iconSize="lg">Dashboard</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-person" iconSize="lg">Profiles</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-box" iconSize="lg">Components</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-chat-square-text" iconSize="lg">Responses</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-arrow-down-square" iconSize="lg">Requests</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-clouds" iconSize="lg">Instances</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-database" iconSize="lg">Databases</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-filter" iconSize="lg">Filters</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/">
                            <CDBSidebarMenuItem icon="bi bi-plus-square" iconSize="lg">Composes</CDBSidebarMenuItem>
                        </NavLink>
                        <NavLink to="/" className="activeClicked">
                            <CDBSidebarMenuItem icon="bi bi-spellcheck" iconSize="lg">Grammars</CDBSidebarMenuItem>
                        </NavLink>
                    </CDBSidebarMenu>
                </CDBSidebarContent>

                <CDBSidebarFooter style={{ textAlign: 'center' }}>
                    <a href="https://steve.care"  class="logo">
                        <img src="assets/logo.svg" width="151" height="192" alt="Steve Care Software" />
                        <span>Steve Care</span>
                    </a>
                </CDBSidebarFooter>
        </CDBSidebar>
    </div>
    );
  }
}
