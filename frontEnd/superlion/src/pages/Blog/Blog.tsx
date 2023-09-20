import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
const useStyles = makeStyles((theme: Theme) => ({
    root: {

    },
}))

const Blog = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                blog
            </Box>
        </>
    )
}

export default Blog;