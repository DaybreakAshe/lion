import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
    },
}))

const SetBlogInfo = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
               
            </Box>
        </>
    )
}

export default SetBlogInfo;

