import { Box, Theme, TextField, Button } from "@mui/material";
import { makeStyles } from '@mui/styles'
import homeImage from "../../../src/assets/images/home/index.jpg"
import { useState } from "react";
const useStyles = makeStyles((theme: Theme) => ({
    content: {
        width: "100%",
        height: "100%",
    },
    imageBox: {
        width: "100%",
        height: "100vh",
        backgroundImage: `url(${homeImage})`,
        backgroundSize: "cover",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
    searchBox: {
        width: "600px",
        height: "300px",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "rgba(255,255,255,.5)",
        backdropFilter: "blur(10px)",
        borderRadius: "10px",
        padding: "20px",
    },
    inputStyle: {
        width: "400px",
        "& .css-9ddj71-MuiInputBase-root-MuiOutlinedInput-root":{
            borderRadius: "30px",
        }
    },
    buttonStyle: {
        width: "120px",
        height: "55px",
    }
}))

const Main = () => {
    const classes = useStyles()
    const [searchValue, setSearchValue] = useState("")
    return (
        <>
            <Box className={classes.content}>
                <Box className={classes.imageBox}>
                    <Box className={classes.searchBox}>
                        <TextField
                            value={searchValue}
                            type="text"
                            onChange={(e) => setSearchValue(e.target.value)}
                            className={classes.inputStyle}
                        />
                        <div style={{width:"20px"}}></div>
                        <Button
                            variant="contained"
                            className={classes.buttonStyle}
                            sx={{
                                borderRadius: "30px",
                            }}
                        >
                            Search
                        </Button>
                    </Box>
                </Box>
            </Box>
        </>
    )
}

export default Main