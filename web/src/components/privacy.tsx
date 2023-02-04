import { Dialog, Typography } from "@mui/material";

export default function Privacy(props: {
  open: boolean;
  handleClose: () => void;
}) {
  const { open, handleClose } = props;
  return (
    <Dialog open={open} onClose={handleClose}>
      <div style={{ padding: "30px" }}>
        <Typography>
          We will collect and use your personal information (which may include
          cookies we collect through your use of realestate.com.au and our other
          websites) to give you a personalised user experience (eg. if you’re
          advertising a property, we may recommend to you products to optimise
          your campaign performance, or, if you’re searching for a property, we
          may recommend to you properties you may be interested in). We may also
          contact you to promote our services or those of third parties. Our
          Privacy Policy further explains how we collect, use and disclose
          personal information and how to access, correct or complain about the
          handling of personal information.
        </Typography>
        <Typography
          sx={{
            textAlign: "center",
            fontSize: "600",
            marginTop: "20px",
            color: "rgb(1,98,157)",
            cursor: "pointer",
            fontWeight: "600",
            "&:hover": {
              textDecoration: "underline",
            },
          }}
          onClick={handleClose}
        >
          Close
        </Typography>
      </div>
    </Dialog>
  );
}
