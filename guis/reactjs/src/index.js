import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router } from 'react-router-dom';

import { Body } from './layouts/body/index.js';

// fonts:
import './assets/fonts/Megrim-Regular.ttf';
import './assets/fonts/Nunito-Italic-VariableFont_wght.ttf';
import './assets/fonts/Nunito-VariableFont_wght.ttf';

// css:
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap-icons/font/bootstrap-icons.css';
import './index.css';




class Main extends React.Component {
  render() {
    return (
        <Router>
            <div className="App">
                <Body />
            </div>
        </Router>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<Main />);
