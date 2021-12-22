function DoomModal(props) {
  const {classes} = props;
  const doomUrl = process.env.REACT_APP_DOOM_URL ? `https://${process.env.REACT_APP_DOOM_URL}` : "http://localhost:8000"
  return (
    <div className={classes.modal} id="doom-container">
        <iframe className={classes.iframe} scrolling="no" src={doomUrl} title="DOOM WASM"></iframe>
    </div>
  )
}

export default DoomModal;