import { createGlobalStyle } from 'styled-components';

export const GlobalStyle = createGlobalStyle`
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }
  
  html,
  body {
    margin: 0;
    padding: 0;
    font-size: 1.1rem;
    font-family: system-ui, Helvetica, Sans-Serif;
  }

  body {
    background: #c2e3f9;
    background: -webkit-linear-gradient(top left, #c2e3f9, #6262db);
    background: -moz-linear-gradient(top left, #c2e3f9, #6262db);
    background: linear-gradient(to bottom right, #c2e3f9, #6262db);
  }

`;
