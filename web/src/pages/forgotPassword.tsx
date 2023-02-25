import { EmailOutlined } from "@mui/icons-material";
import {
  Alert,
  Button,
  Card,
  Divider,
  InputAdornment,
  Link,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import Privacy from "../components/privacy";
import RentHomeLogo from "../components/rentHomeLogo";
import { API_ADDRESS } from "../config";
import { saveResetPasswordTokenInLocalStorage } from "../helpers/auth";
import "../styles/index.css";
import { ErrorResponse, ForgotPasswordResponse } from "../types/types";

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
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const onSubmit: SubmitHandler<IForgotPasswordInput> = async (input) => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/forgot-password`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(input),
      });
      if (res.ok) {
        const data: ForgotPasswordResponse = await res.json();
        saveResetPasswordTokenInLocalStorage(data.reset_token);
        navigate("/confirm-forgot-password");
      } else {
        const data: ErrorResponse = await res.json();
        setError(data.message);
      }
    } catch (error) {
      setError("Something went wrong, please try again.");
    }
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
          {error && <Alert severity="warning">{error}</Alert>}
          <TextField
            variant="outlined"
            placeholder="Email address"
            {...register("email", {
              required: {
                value: true,
                message: "Please enter a valid email address.",
              },
            })}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <EmailOutlined />
                </InputAdornment>
              ),
            }}
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
