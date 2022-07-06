import React from 'react';
import styled from 'styled-components';

interface LogoProps {}

const HeaderContainer = styled.div`
	margin: 0 auto;
	width: 90%;
	min-height: 5rem;
	display: flex;
	align-items: center;

	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		justify-content: center;
	}
`;

const Wrapper = styled.header`
	width: 100%;
`;

const LogoText = styled.h1`
	color: ${({ theme }) => theme.colors.secondary};
	letter-spacing: 0.4rem;
`;

const Header: React.FC<LogoProps> = () => {
	return (
		<Wrapper>
			<HeaderContainer>
				<LogoText>SHORTURL</LogoText>
			</HeaderContainer>
		</Wrapper>
	);
};
export default Header;
