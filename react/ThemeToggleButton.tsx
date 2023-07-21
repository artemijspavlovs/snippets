"use client";

import { DarkModeTwoTone, LightModeTwoTone } from "@mui/icons-material";
import { IconButton, useTheme } from "@mui/material";
import { createContext, useContext } from "react";

export const ColorThemeContext = createContext({ toggleColorMode: () => {} });

export const ToggleColorThemeButton = () => {
  const theme = useTheme();
  const colorMode = useContext(ColorThemeContext);

  return (
    <>
      <IconButton
        sx={{ ml: 1 }}
        onClick={colorMode.toggleColorMode}
        color="inherit"
      >
        {theme.palette.mode === "dark" ? (
          <LightModeTwoTone />
        ) : (
          <DarkModeTwoTone />
        )}
      </IconButton>
    </>
  );
};
