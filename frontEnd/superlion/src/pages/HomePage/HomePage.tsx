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
                #111
            </Box>
        </>
    )
}

export default HomePage;