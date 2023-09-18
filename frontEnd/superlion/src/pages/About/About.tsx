import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
const useStyles = makeStyles((theme: Theme) => ({
    root: {

    },
}))

const About = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                about
            </Box>
        </>
    )
}

export default About;