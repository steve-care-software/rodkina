import React from 'react';

import { TokenCreate } from './token_create/index.js';

export class Application extends React.Component {
  render() {
    return (
        <div id="application">
            <TokenCreate />
        </div>
    );
  }
}
