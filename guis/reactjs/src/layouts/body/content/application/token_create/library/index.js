import React from 'react';
import Accordion from 'react-bootstrap/Accordion';

import {
  CDBSidebar,
  CDBSidebarContent,
  CDBSidebarHeader,
} from 'cdbreact';

import { Channels } from './channels/index.js';
import { Composes } from './composes/index.js';
import { Exceptions } from './exceptions/index.js';
import { ExternalGrammars } from './externals/index.js';
import { Values } from './values/index.js';

import './index.css';

export class Library extends React.Component {
  render() {
    return (
        <div id="library">
            <CDBSidebar minWidth="80px" maxWidth="400px" textColor="#fff" backgroundColor="#333">
                <CDBSidebarHeader prefix={<i className="fa fa-bars fa-large"></i>}>
                    <a href="/" className="text-decoration-none" style={{ color: 'inherit' }}>
                        Library
                    </a>
                </CDBSidebarHeader>

                <CDBSidebarContent className="sidebar-content">
                    <Accordion defaultActiveKey="0">
                          <Accordion.Item eventKey="0">
                              <Accordion.Header>Channels</Accordion.Header>
                              <Accordion.Body>
                                  <Channels />
                              </Accordion.Body>
                          </Accordion.Item>
                          <Accordion.Item eventKey="1">
                              <Accordion.Header>External Grammars</Accordion.Header>
                              <Accordion.Body>
                                  <ExternalGrammars />
                              </Accordion.Body>
                          </Accordion.Item>
                          <Accordion.Item eventKey="2">
                              <Accordion.Header>Exceptions</Accordion.Header>
                              <Accordion.Body>
                                  <Exceptions />
                              </Accordion.Body>
                          </Accordion.Item>
                          <Accordion.Item eventKey="3">
                              <Accordion.Header>Composes</Accordion.Header>
                              <Accordion.Body>
                                  <Composes />
                              </Accordion.Body>
                          </Accordion.Item>
                          <Accordion.Item eventKey="4">
                              <Accordion.Header>Values</Accordion.Header>
                              <Accordion.Body>
                                  <Values />
                              </Accordion.Body>
                          </Accordion.Item>
                  </Accordion>
                </CDBSidebarContent>
        </CDBSidebar>
    </div>
    );
  }
}
