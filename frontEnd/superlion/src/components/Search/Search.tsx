import { Box, Theme, TextField } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { useState } from "react";
import SearchIcon from '@mui/icons-material/Search';
const useStyles = makeStyles((theme: Theme) => ({
    root: {

    },
    searchBox: {
        height: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        marginRight: "30px",
    },
    inputStyle: {
        width: "400px",
        "& .css-9ddj71-MuiInputBase-root-MuiOutlinedInput-root": {
            borderRadius: "30px",
        }
    },
    searchIcon: {
        color: "#1a73e8",
        marginLeft: "15px",
        fontSize: "30px",
        cursor: "pointer",
    }
}))

const Search = () => {
    const classes = useStyles()
    const [searchValue, setSearchValue] = useState("")
    const search = () => {

    }
    return (
        <>
            <Box className={classes.root}>
                <Box className={classes.searchBox}>
                    <TextField
                        size="small"
                        value={searchValue}
                        type="text"
                        onChange={(e) => setSearchValue(e.target.value)}
                        className={classes.inputStyle}
                        placeholder="Search"
                        onKeyDown={(e) => {
                            if (e.key === "Enter") {
                                search()
                            }
                        }}
                    />
                    <SearchIcon
                        onClick={search}
                        className={classes.searchIcon}
                    />
                </Box>
            </Box>
        </>
    )
}

export default Search;