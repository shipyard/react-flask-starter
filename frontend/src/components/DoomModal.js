function DoomModal(props) {
  const {classes} = props;
  return (
    <div className={classes.modal} id="doom-container">
        <iframe className={classes.iframe} scrolling="no" src="http://localhost:8000" title="DOOM WASM"></iframe>
    </div>
  )
}

export default DoomModal;