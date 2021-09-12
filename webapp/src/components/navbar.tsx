import PersonIcon from "@mui/icons-material/Person";
import { AppBar, IconButton, Toolbar } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";

export function Navbar(): JSX.Element {
  const [cookie] = useCookies(["GAKUJO_SESSION"]);
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          {cookie.GAKUJO_SESSION ? (
            <IconButton>
              <PersonIcon style={{ fill: "black" }} />
            </IconButton>
          ) : (
            <Link to="/signin">Login</Link>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}
