import React, { useEffect, useRef, useState } from 'react';
import styled from 'styled-components';
import { Container } from '../styles/StyledElements';
import Card from './Card';
import Spinner from './Spinner';

interface ConverterProps {}

const Wrapper = styled.section`
	flex-grow: 1;
	margin-top: 3rem;
`;

const MainContainer = styled(Container)`
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: flex-start;
	height: 100%;
`;

const Button = styled.button<{ copied?: boolean }>`
	border: none;
	outline: none;
	color: white;
	background-color: ${({ theme }) => theme.colors.secondary};
	padding: 0.3em 1em;
	font-size: 1.1rem;
	border-radius: 0.5rem;
	margin-left: 2rem;
	&:hover {
		cursor: pointer;
		transform: scale(1.01);
	}

	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		margin-left: 0;
		margin-top: 0.7rem;
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
	}
`;

const CopyButton = styled(Button)`
	margin-left: 0.5rem;
	background-color: ${({ copied, theme }) =>
		copied ? theme.colors.lightPrimary : theme.colors.secondary};
	color: ${({ copied, theme }) => (copied ? 'black' : 'white')};
	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		margin-top: 0.2rem;
		width: 100%;
	}
`;

const Input = styled.input<{ valid?: boolean }>`
	font-size: inherit;
	font-family: inherit;
	padding: 0.3em 0.5em;
	border: 1px solid ${({ valid }) => (valid ? 'grey' : 'red')};
	border-radius: 0.5rem;
	text-align: center;
	flex-grow: 1;
	&:focus {
		border: 1px solid grey;
		outline: none;
	}
	flex-grow: 1;
	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		width: 100%;
		margin: 0 auto;
	}
`;

const InputContainer = styled.div`
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;

	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		width: 100%;
		flex-direction: column;
	}
`;

const ShortURL = styled.div`
	width: 100%;
	margin-top: 1rem;
`;

const Error = styled.p`
	font-size: inherit;
	color: red;
	margin-top: 0.1rem;
	margin-left: 0.6rem;
	@media (max-width: ${({ theme }) => theme.breakpoints.small}) {
		margin-left: 0;
	}
`;

const Converter: React.FC<ConverterProps> = () => {
	const [loading, setLoading] = useState(false);
	const [encodedURL, setEncodedURL] = useState('');
	const [error, setError] = useState('');
	const [copied, setCopied] = useState(false);
	const inputRef = useRef<HTMLInputElement>(null);
	const outputRef = useRef<HTMLInputElement>(null);

	useEffect(() => {
		if (inputRef == null || inputRef.current == null) return;
		inputRef.current.focus();
	}, []);

	const copyHandler = () => {
		if (outputRef == null || outputRef.current == null) return;
		navigator.clipboard.writeText(outputRef.current.value);
		setCopied(true);
	};

	const clickHandler = async () => {
		if (inputRef == null || inputRef.current == null) return;
		if (inputRef.current.value == '') {
			setError('Enter the longURL');
			return;
		}
		setLoading(true);
		setCopied(false);
		const apiresponse = await fetch('/encode', {
			body: JSON.stringify({
				url: inputRef.current.value,
			}),
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
		});
		const responsedata = await apiresponse.json();
		if (!apiresponse.ok) {
			setError(responsedata.message);
		} else {
			setEncodedURL(window.location.href + responsedata.encodedURL);
		}
		setLoading(false);
	};

	return (
		<Wrapper>
			<MainContainer>
				<Card>
					<InputContainer>
						<Input
							spellCheck={'false'}
							ref={inputRef}
							placeholder="Enter a long URL"
							onChange={() => {
								setError('');
							}}
							valid={error == ''}
						/>
						<Button onClick={clickHandler}>
							{!loading && (
								<>
									{encodedURL && <>Shorten another!</>}
									{!encodedURL && <>Shorten!</>}
								</>
							)}
							{loading && <Spinner />}
						</Button>
					</InputContainer>
					{error && <Error>{error}</Error>}
					{encodedURL && (
						<ShortURL>
							<InputContainer>
								<Input
									spellCheck={'false'}
									ref={outputRef}
									placeholder="Enter a long URL"
									contentEditable="false"
									value={encodedURL}
									readOnly
									valid
								/>
								<CopyButton
									copied={copied}
									onClick={copyHandler}
								>
									{copied && <>Copied!</>}
									{!copied && <>Copy</>}
								</CopyButton>
							</InputContainer>
						</ShortURL>
					)}
				</Card>
			</MainContainer>
		</Wrapper>
	);
};
export default Converter;
