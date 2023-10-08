import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import MenuItem from "./MenuItem.tsx";
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
    },
}))

const Menu = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                <MenuItem title="Blog" path="/blog" />
                <MenuItem title="About" path="/about" />
                <MenuItem title="Contacts" path="/contacts" />
            </Box>
        </>
    )
}

export default Menu;