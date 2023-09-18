import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
const useStyles = makeStyles((theme: Theme) => ({
    root: {

    },
}))

const HomePage = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                <img src="https://z1.ax1x.com/2023/09/18/pP4VWo8.jpg" alt="" style={{
                    width: "100%",
                    height: "100%",
                }}/>
            </Box>
        </>
    )
}

export default HomePage;