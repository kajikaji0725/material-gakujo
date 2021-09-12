import { Menu, MenuItem } from "@mui/material";
import React, { useState } from "react";

export default function UserMenu(props: {
    anchorEl: HTMLElement | null,
    onClose: () => void,
}): JSX.Element {
  const open = Boolean(props.anchorEl);
  const handleClose = () => props.onClose();

  return (
     <Menu
        anchorEl={props.anchorEl}
        open={open}
        onClose={handleClose}
     >
         <MenuItem onClick={handleClose}></MenuItem>
     </Menu>
  )
}
