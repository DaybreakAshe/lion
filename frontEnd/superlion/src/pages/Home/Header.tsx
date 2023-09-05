import { Box, Theme, TextField } from "@mui/material";
import { makeStyles } from '@mui/styles'
import logo from "../../../src/assets/images/home/logo.png"
import { useState } from "react";
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
    },
    searchBox: {
        width: "600px",
        height: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
    inputStyle: {
        width: "400px",
        "& .css-9ddj71-MuiInputBase-root-MuiOutlinedInput-root": {
            borderRadius: "30px",
        }
    },
    buttonStyle: {
        width: "100px",
        height: "40px",
    }
}))

const Header = () => {
    const classes = useStyles()
    const [searchValue, setSearchValue] = useState("")

    return (
        <>
            <Box className={classes.content}>
                <img src={logo} alt="logo" style={{
                    width: "60px",
                    height: "60px",
                }} />
                <Box className={classes.searchBox}>
                    <TextField
                        size="small"
                        value={searchValue}
                        type="text"
                        onChange={(e) => setSearchValue(e.target.value)}
                        className={classes.inputStyle}
                    />
                    {/* <SearchIcon style={{marginLeft:"20px"}}/> */}
                </Box>
                <div></div>
            </Box>
        </>
    )
}

export default Header