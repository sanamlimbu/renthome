import {
  Button,
  Card,
  Divider,
  Link,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import Privacy from "../components/privacy";
import RentHomeLogo from "../components/rentHomeLogo";
import "../styles/index.css";

interface IForgotPasswordInput {
  email: string;
}

export default function ForgotPasswordPage() {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<IForgotPasswordInput>();
  const [open, setOpen] = useState(false);

  const onSubmit: SubmitHandler<IForgotPasswordInput> = (data) => {
    console.log(data);
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        <Typography textAlign={"center"} fontWeight={"700"} fontSize={"18px"}>
          Forgot your password?
        </Typography>
        <Typography textAlign={"center"}>
          Enter your email and we’ll send you a code you can use to update your
          password.
        </Typography>
        <form className="Form" onSubmit={handleSubmit(onSubmit)}>
          <TextField
            variant="outlined"
            placeholder="Email address"
            {...register("email", {
              required: {
                value: true,
                message: "Please enter a valid email address.",
              },
            })}
          />
          {errors.email && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.email.message}
            </span>
          )}
          <Button
            variant="contained"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Reset my password
          </Button>
        </form>
        <Link
          className="Link"
          href="/login"
          sx={{ paddingTop: "20px", paddingBottom: "20px" }}
        >
          Go back to sign in.
        </Link>
        <Divider />
        <Typography
          sx={{ textAlign: "center", fontSize: "12px", cursor: "pointer" }}
          onClick={() => setOpen(true)}
        >
          Personal Information Collection Statement
        </Typography>
      </Card>
      <Privacy open={open} handleClose={() => setOpen(false)} />
    </div>
  );
}
