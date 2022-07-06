import React from 'react';
import Converter from './components/Converter';
import Header from './components/Header';
import styled from 'styled-components';

interface AppProps {}

const Wrapper = styled.main`
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	width: 100%;
	height: 100vh;
`;

const App: React.FC<AppProps> = ({}) => {
	return (
		<Wrapper>
			<Header />
			<Converter />
		</Wrapper>
	);
};
export default App;
