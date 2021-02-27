import { useEffect, useState } from 'react';
import {
  Box,
  Button,
  Card,
  GridList,
  GridListTile,
  Link,
} from '@material-ui/core';

import localStackLogo from '../assets/images/localStackLogo.png';

function UploadCard(props) {
  const [isLoaded, setIsLoaded] = useState(false);
  const [data, setData] = useState([]);
  const [error, setError] = useState(null);
  const { classes } = props;

  const fetchBucketFiles = () => {
    fetch('api/v1/files')
      .then((res) => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          const bucket = result['Name'];
          const filenames = result['Contents']
            ? result['Contents']
                .map((file) => `${bucket}/${file['Key']}`)
                .reverse()
            : [];
          setData(filenames);
        },
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      );
  };

  useEffect(() => fetchBucketFiles(), []);

  const submitFile = async (e) => {
    e.preventDefault();
    const file = e.target.files;
    try {
      if (!file) {
        throw new Error('Please select an image');
      }
      const formData = new FormData();
      formData.append('file', file[0]);
      await fetch('/api/v1/files/upload/', {
        method: 'POST',
        body: formData,
      });
      fetchBucketFiles();
    } catch (error) {
      alert(error);
    }
  };

  return (
    <Card className={classes.card}>
      <h2>Store files in S3, locally</h2>
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
        <img
          width="75px"
          height="100%"
          src={localStackLogo}
          alt="LocalStack Logo"
        />
      </Box>
      <p>Click below to select and upload an image from your computer.</p>
      <p>
        The image will be <b>stored in a local S3 bucket</b>, powered by{' '}
        <Link
          className={classes.link}
          target="_blank"
          rel="noopener"
          href="https://github.com/localstack/localstack"
        >
          <b>LocalStack</b>
        </Link>{' '}
        - a fully functional local AWS cloud stack, and displayed below.
      </p>
      <form>
        <Box
          display="flex"
          flexDirection="column"
          alignItems="center"
          justifyContent="center"
          mt={2}
        >
          <Button
            variant="contained"
            className={classes.contained}
            component="label"
            size="small"
          >
            Upload Image
            <input
              type="file"
              accept=".jpg,.jpeg,.png,.gif"
              onChange={(event) => {
                submitFile(event);
              }}
              hidden
            />
          </Button>
        </Box>
      </form>

      {data ? (
        <Box display="flex" flexDirection="row" justifyContent="center">
          <div className={classes.gridContainer}>
            <GridList cellHeight={100} cols={3}>
              {data.map((src) => (
                <GridListTile key={src} cols={1}>
                  <img src={src} alt="Uploaded to LocalStack" />
                </GridListTile>
              ))}
            </GridList>
          </div>
        </Box>
      ) : (
        ''
      )}
    </Card>
  );
}

export default UploadCard;
