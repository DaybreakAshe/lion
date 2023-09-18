import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import Header from "../header/Header";
import Main from "../../router/Main";

const useStyles = makeStyles((_theme: Theme) => ({
    root: {
        position: "relative",
    }
}))

const Home = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                <Header />
                <Main />
            </Box>
        </>
    )
}
export default Home