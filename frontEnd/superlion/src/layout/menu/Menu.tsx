import { Box, Theme } from "@mui/material";
import { makeStyles } from "@mui/styles";
import MenuItem from "./MenuItem.tsx";
const useStyles = makeStyles((_theme: Theme) => ({
  root: {},
}));

const Menu = () => {
  const classes = useStyles();
  return (
    <>
      <Box className={classes.root}>
        <MenuItem title="博客" path="/blog" />
        <MenuItem title="关于" path="/about" />
      </Box>
    </>
  );
};

export default Menu;
