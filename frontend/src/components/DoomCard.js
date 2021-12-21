import { Card } from '@material-ui/core';

import doomLogo from '../assets/images/doom.png';
function DoomCard(props) {

	const { classes, handleClick } = props;

	return (
		<Card className={classes.card}>
			<h2 className={classes.cardHeader}>Play DOOM</h2>
			<p>
				Play DOOM embedded on this web page.
			</p>
      <div className={classes.doomButtonContainer}>
        <img id="doom-button" src={doomLogo} onClick={handleClick}/>
      </div>
		</Card>
	);
}

export default DoomCard;
