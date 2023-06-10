import { createTheme } from "@mui/material";

declare module "@mui/material/styles" {
  interface PaletteOptions {
    lightgrey: string;
    dimblack: string;
  }
}
const theme = createTheme({
  palette: {
    secondary: {
      main: "#FFEC3D",
    },
    success: {
      main: "#44b700",
    },
    background: {
      default: "#F7F9FC",
    },
    lightgrey: "#f7f7f7",
    dimblack: "#1a1a1a82",
    common: {
      black: "#1a1a1a",
    },
  },
  breakpoints: {
    values: {
      xs: 0,
      sm: 600,
      md: 960,
      lg: 1280,
      xl: 1440,
    },
  },
  typography: {
    fontSize: 16,
    fontWeightLight: 300,
    fontWeightRegular: 400,
    fontWeightMedium: 600,
    fontWeightBold: 700,
    h1: {
      fontSize: "2rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    h2: {
      fontSize: "1.75rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    h3: {
      fontSize: "1.5rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    h4: {
      fontSize: "1.25rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    h5: {
      fontSize: "1.125rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    h6: {
      fontSize: "1.0625rem",
      fontWeight: 600,
      lineHeight: 1.2,
    },
    body1: {
      fontSize: 16,
    },
    button: {
      textTransform: "none",
      fontWeight: 500,
    },
  },
  components: {
    MuiButton: {
      defaultProps: {
        color: "inherit",
      },
    },
  },
});

export default theme;
