import React, { ReactNode } from 'react';
import styled from 'styled-components';

interface CardProps {
	children: ReactNode;
}

const Wrapper = styled.div`
	min-width: 50%;
	border-radius: 1rem;
	/* min-height: 15rem; */
	padding: 2em;
	background-color: white;

	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		min-width: 90%;
		margin: 0 auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
	}
`;

const Card: React.FC<CardProps> = ({ children }) => {
	return <Wrapper>{children}</Wrapper>;
};
export default Card;
