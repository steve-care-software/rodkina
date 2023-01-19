import React from 'react';
import Container from 'react-bootstrap/Container';

import { BreadcrumbMenu } from './breadcrumb/index.js';
import { Application } from './application/index.js';

import './index.css';

export class Content extends React.Component {
  render() {
    return (
        <Container id="content" fluid>
            <BreadcrumbMenu />
            <Application />
        </Container>
    );
  }
}
