import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router } from 'react-router-dom';

import { TopMenu } from './layouts/topmenu/index.js';
import { Body } from './layouts/body/index.js';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap-icons/font/bootstrap-icons.css';
import './index.css';


class Main extends React.Component {
  render() {
    return (
        <Router>
            <div className="App">
                <TopMenu />
                <Body />
            </div>
        </Router>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<Main />);
