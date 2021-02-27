import { Card, Box, Link } from '@material-ui/core';

import logo from '../assets/images/logo.png';

function DeployCard(props) {
  const { classes } = props;

  return (
    <Card className={classes.card}>
      <h2>Deploy your app to the cloud</h2>
      <Box
        display="flex"
        flexDirection="row"
        alignItems="center"
        mt={-3}
        mb={-1}
      >
        <Box mr={1}>
          <h3>Powered by</h3>
        </Box>
        <img width="120px" height="100%" src={logo} alt="Shipyard Logo" />
      </Box>
      <p>
        Use{' '}
        <Link
          className={classes.link}
          target="_blank"
          rel="noopener"
          href="https://shipyard.build"
        >
          <b>Shipyard</b>
        </Link>{' '}
        to deploy your app to the cloud, following these simple steps:
      </p>
      <ul className={classes.list}>
        <li>
          Go to{' '}
          <Link
            className={classes.link}
            target="_blank"
            rel="noopener"
            href="https://shipyard.build"
          >
            <b>https://shipyard.build</b>
          </Link>
          .
        </li>
        <li>
          <b>Sign in</b> using your{' '}
          <Link
            className={classes.link}
            target="_blank"
            rel="noopener"
            href="https://github.com"
          >
            <b>GitHub</b>
          </Link>{' '}
          credentials.
        </li>
        <li>
          Click <b>Add Repository</b> to select the repo for this starter
          project.
        </li>
        <li>
          <b>Select the branch</b> you'd like to deploy.
        </li>
        <li>
          Click <b>Deploy</b> to deploy the app!
        </li>
      </ul>
      <p>
        Once the app is up and running, you'll be able to see it by clicking the{' '}
        <b>View Live Environment</b> icon.
      </p>
    </Card>
  );
}

export default DeployCard;
