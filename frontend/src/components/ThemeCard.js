import { Card, Link } from '@material-ui/core';

function ThemeCard(props) {
	const codeBlock_1 = `94 |  const theme = createMuiTheme({
95 |    palette: {
96 |      type: 'dark',
97 |    },
98 |  });`;

	const codeBlock_2 = `94 |  const theme = createMuiTheme({
95 |    palette: {
96 |      type: 'light',
97 |    },
98 |  });`;

	const { classes } = props;

	return (
		<Card className={props.classes.card}>
			<h2 className={props.classes.cardHeader}>Update appearance</h2>
			<p>
				Modify the appearance of your app by using Material-UI's built-in
				themes, or by creating a custom theme.
			</p>
			<p>
				To switch from <b>Dark Mode</b> to <b>Light Mode</b>, open{' '}
				<code>
					<Link
						className={classes.link}
						href={`${process.env.REACT_APP_STARTER_REPO_URL}frontend/src/components/App.js#L94-L98`}
						target="_blank"
					>
						frontend/src/components/App.js
					</Link>
				</code>{' '}
				and replace the following code:
			</p>
			<pre>
				<code className="language-js">{codeBlock_1}</code>
			</pre>
			<p>with:</p>
			<pre>
				<code className="language-js">{codeBlock_2}</code>
			</pre>
			<p>
				<Link
					className={classes.link}
					href="https://material-ui.com/customization/theming/"
					target="_blank"
				>
					<b>Click here</b>
				</Link>{' '}
				to learn more about Material-UI themes.
			</p>
		</Card>
	);
}

export default ThemeCard;
