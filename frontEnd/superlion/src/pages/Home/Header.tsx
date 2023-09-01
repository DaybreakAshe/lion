import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import logo from "../../../src/assets/images/home/logo.png"
const useStyles = makeStyles((theme: Theme) => ({
    content: {
        width: "100%",
        height: "65px",
        backgroundColor: "#fff",
        position: "fixed",
        top: "0",
        left: "0",
        zIndex: "1000",
        boxShadow: "0 0 10px rgba(0,0,0,.1)",
        padding: "0 20px",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
    },
    logoBox: {
        width: "60px",
        height: "60px",
    }
}))

const Header = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.content}>
                <img src={logo} alt="logo" style={{
                    width: "60px",
                    height: "60px",
                }} />
            </Box>
        </>
    )
}

export default Header