import {
  Card,
  Chip,
  Link,
  List,
  ListItem,
  ListItemText,
} from '@material-ui/core';

function MaterialCard(props) {
  const { classes } = props;
  const seasons = ['Winter', 'Spring', 'Summer', 'Fall'];
  const codeBlock_1 = `27 | {seasons.map((season) => (
28 |    <Chip key={season} label={season} />
29 |  ))} `;
  const codeBlock_2 = `27 | {seasons.map((season) => (
28 |    <ListItem button>
29 |      <ListItemText primary={season} />
30 |    </ListItem>
31 |  ))} `;

  return (
    <Card className={classes.card}>
      <h2 className={classes.cardHeader}>Use Material-UI components</h2>
      <p>Material-UI provides a variety of UI components to render data.</p>
      <List>
        {seasons.map((season) => (
          <Chip key={season} label={season} />
        ))}
      </List>
      <p>
        To change the <b>chips</b> above to <b>list items</b>, open{' '}
        <code>
          <Link
            className={classes.link}
            href={`${process.env.REACT_APP_STARTER_REPO_URL}frontend/src/components/MaterialCard.js#L27-L29`}
            target="_blank"
          >
            frontend/src/components/MaterialCard.js
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
        Read the Material-UI docs to learn more about components:{' '}
        <Link
          className={classes.link}
          href="https://material-ui.com"
          target="_blank"
          rel="noopener"
        >
          <b>https://material-ui.com</b>
        </Link>
        .
      </p>
    </Card>
  );
}

export default MaterialCard;
