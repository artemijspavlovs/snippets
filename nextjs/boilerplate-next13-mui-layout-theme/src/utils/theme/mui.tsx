"use client";

import {
  createTheme,
  PaletteMode,
  PaletteOptions,
  Typography,
} from "@mui/material";
import { TypographyOptions } from "@mui/material/styles/createTypography";

const headerFont = "OutfitVariable, sans-serif";
const textFont = "InconsolataVariable, monoscope";

const typographyOptions: TypographyOptions = {
  h1: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  h2: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  h3: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  h4: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  h5: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  h6: {
    fontFamily: headerFont,
    fontWeight: 600,
  },
  body1: {
    fontFamily: textFont,
    fontWeight: 400,
  },
  body2: {
    fontFamily: textFont,
    fontWeight: 400,
  },
  subtitle1: {
    fontFamily: headerFont,
    fontWeight: 400,
  },
  subtitle2: {
    fontFamily: headerFont,
    fontWeight: 400,
  },
  button: {
    fontFamily: headerFont,
    fontWeight: 500,
  },
  caption: {
    fontFamily: headerFont,
    fontWeight: 400,
  },
  overline: {
    fontFamily: headerFont,
    fontWeight: 400,
  },
};
// https://bareynol.github.io/mui-theme-creator/#Accordion
const lightPalette: PaletteOptions = {
  mode: "light",
  primary: {
    main: "#293241",
  },
  secondary: {
    main: "#F75C03",
  },
  background: {
    default: "#F4F1DE",
    paper: "#D0C6AC",
  },
  text: {
    primary: "#293241",
    secondary: "#6D6875",
    disabled: "rgba(156,3,3,0.38)",
  },
  error: {
    main: "#EF476F",
  },
  warning: {
    main: "#FFD166",
  },
  info: {
    main: "#118AB2",
  },
  success: {
    main: "#06D6A0",
  },
  divider: "#F4F1DE",
};
const darkPalette: PaletteOptions = {
  mode: "dark",
  primary: {
    main: "#F4F1DE",
  },
  secondary: {
    main: "#FF8E3E",
  },
  background: {
    default: "#293241",
    paper: "#525B6C",
  },
  text: {
    primary: "#F4F1DE",
    secondary: "#D0C6AC",
    disabled: "rgba(82,91,108,0.38)",
  },
  error: {
    main: "#EF476F",
  },
  warning: {
    main: "#FFD166",
  },
  info: {
    main: "#5ABAE4",
  },
  success: {
    main: "#06D6A0",
  },
  divider: "#F4F1DE",
};

export const WNThemeGenerate = (mode: PaletteMode) => ({
  palette: {
    mode,
    ...(mode === "light" ? lightPalette : darkPalette),
  },
  typography: typographyOptions,
});
