import Prism from 'prismjs';
import 'prismjs/themes/prism-tomorrow.css';
import { Grid, Box, Link, Typography } from '@material-ui/core';
import {
  makeStyles,
  createMuiTheme,
  ThemeProvider,
} from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import ConnectCard from './ConnectCard';
import DeployCard from './DeployCard';
import MaterialCard from './MaterialCard';
import ThemeCard from './ThemeCard';
import UploadCard from './UploadCard';

import logo from '../assets/images/logo.png';

const useStyles = makeStyles((theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    padding: '30px',
    width: 'calc(100% - 16px)',
    marginLeft: '8px',
  },
  card: {
    height: '700px',
    padding: '10px 25px',
    overflow: 'auto',
  },
  cardHeader: {
    marginBottom: '30px',
  },
  img: {
    width: '210px',
    background: 'white',
    padding: '0 10px',
    height: '100%',
    margin: '0 20px',
  },
  box: {
    flexDirection: 'row',
    alignItems: 'center',
    [theme.breakpoints.down('sm')]: {
      flexDirection: 'column',
      textAlign: 'center',
    },
  },
  gridContainer: {
    marginTop: 20,
    display: 'flex',
    flexWrap: 'wrap',
    width: 500,
  },
  form: {
    marginTop: '50',
  },
  response: {
    fontFamily: 'monospace',
    color: 'limeGreen',
    fontSize: '1.3em',
    background: 'black',
    padding: '0px 5px',
  },
  list: {
    fontSize: '1.1em',
    '& li': {
      marginBottom: '8px',
    },
  },
  contained: {
    color: 'white',
    fontWeight: '600',
    background: '#787878',
    '&:hover': {
      background: '#888',
    },
    borderRadius: 5,
    margin: '15px 0px',
  },
  pre: {
    whiteSpace: 'pre-wrap',
  },
  link: {
    color: '#35baf6',
  },
}));

function App() {
  const classes = useStyles();

  const theme = createMuiTheme({
    palette: {
      type: 'dark',
    },
  });

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box>
        <Box className={classes.box} display="flex">
          <Box
            mt="20px"
            display="flex"
            flexDirection="column"
            alignItems="center"
          >
            <img className={classes.img} src={logo} alt="Shipyard logo" />
            <Typography variant="caption">
              Replace this with your logo:
            </Typography>
            <Typography variant="caption">
              <code>'/frontend/src/assets/images/logo.png'</code>{' '}
            </Typography>
          </Box>
          <Box mt={-2}>
            <h1>
              <Link
                className={classes.link}
                target="_blank"
                rel="noopener"
                href="https://github.com/facebook/react"
              >
                React
              </Link>
              -
              <Link
                className={classes.link}
                target="_blank"
                rel="noopener"
                href="https://github.com/pallets/flask"
              >
                Flask
              </Link>
              -
              <Link
                className={classes.link}
                target="_blank"
                rel="noopener"
                href="https://github.com/postgres/postgres"
              >
                Postgres
              </Link>
              -
              <Link
                className={classes.link}
                target="_blank"
                rel="noopener"
                href="https://github.com/localstack/localstack"
              >
                LocalStack
              </Link>{' '}
              Starter Project
            </h1>
          </Box>
        </Box>
        <Grid container className={classes.container} spacing={2}>
          <Grid item xs={12} sm={12} md={6} lg={4}>
            <ThemeCard classes={classes} />
          </Grid>
          <Grid item xs={12} sm={12} md={6} lg={4}>
            <MaterialCard classes={classes} />
          </Grid>
          <Grid item xs={12} sm={12} md={6} lg={4}>
            <ConnectCard classes={classes} />
          </Grid>
          <Grid item xs={12} sm={12} md={6} lg={4}>
            <UploadCard classes={classes} />
          </Grid>
          <Grid item xs={12} sm={12} md={6} lg={4}>
            <DeployCard classes={classes} />
          </Grid>
        </Grid>
      </Box>
    </ThemeProvider>
  );
}

export default App;
