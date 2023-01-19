import React from 'react';

import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';

import './index.css';

export class TopMenu extends React.Component {
  render() {
    return (
        <div id="topmenu">
            <Navbar collapseOnSelect expand="lg" bg="dark">
                <Container fluid>
                    <Navbar.Brand href="/">Steve's Care</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="/">Home</Nav.Link>
                            <NavDropdown title="Projects" id="basic-nav-dropdown">
                                <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
                                <NavDropdown.Item href="#action/3.2">
                                    Another action
                                </NavDropdown.Item>
                                <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
                                <NavDropdown.Divider />
                                <NavDropdown.Item href="#action/3.4">
                                    Separated link
                                </NavDropdown.Item>
                            </NavDropdown>
                            <Nav.Link href="/units">Units</Nav.Link>
                            <Nav.Link href="/blog">Blog</Nav.Link>
                            <Nav.Link href="/newsletter">Newsletter</Nav.Link>
                            <Nav.Link href="/communities">Communities</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                    <Navbar.Collapse className="justify-content-end socialmedia">
                        <Nav.Link href="/"  className="bi bi-github"><span>Github</span></Nav.Link>
                        <Nav.Link href="/"  className="bi bi-twitter"><span>Twitter</span></Nav.Link>
                        <Nav.Link href="/"  className="bi bi-reddit"><span>Reddit</span></Nav.Link>
                        <Nav.Link href="/"  className="bi bi-youtube"><span>Youtube</span></Nav.Link>
                        <Nav.Link href="/"  className="bi bi-facebook"><span>Facebook</span></Nav.Link>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </div>
    );
  }
}
