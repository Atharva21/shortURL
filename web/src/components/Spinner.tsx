import React from 'react';
import styled, { keyframes } from 'styled-components';

const spin = keyframes`
  0%{
    transform: rotate(0deg);
  }
  100%{
    transform: rotate(360deg);
  }
`;

const AnimatedSpinner = styled.div`
	width: 1.5rem;
	height: 1.5rem;
	margin: 0;
	padding: 0;
	border-radius: 50%;
	border: 4px solid transparent;
	border-bottom: 4px solid ${({ theme }) => theme.colors.neutral};
	border-right: 4px solid ${({ theme }) => theme.colors.neutral};
	animation: 750ms ${spin} linear infinite;
`;

const Spinner: React.FC = () => {
	return <AnimatedSpinner />;
};
export default Spinner;
