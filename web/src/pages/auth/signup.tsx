import {
  EmailOutlined,
  LockOutlined,
  Visibility,
  VisibilityOff,
} from "@mui/icons-material";
import {
  Alert,
  Button,
  Divider,
  InputAdornment,
  Link,
  Snackbar,
  TextField,
  Typography,
} from "@mui/material";
import Card from "@mui/material/Card";
import React, { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import Privacy from "../../components/privacy";
import RentHomeLogo from "../../components/rentHomeLogo";
import Social from "../../components/social";
import { API_ADDRESS } from "../../config";
import { socialList } from "../../const";
import "../../styles/index.css";

interface ISignupInput {
  email: string;
  password: string;
}

export default function SignupPage() {
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<ISignupInput>();
  const [open, setOpen] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const onSubmit: SubmitHandler<ISignupInput> = async (input) => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/signup`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(input),
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.message);
        setSnackbarOpen(true);
      } else {
        setSnackbarOpen(false);
        navigate("/");
      }
    } catch (error) {
      setSnackbarOpen(true);
    }
  };

  const handleSnackbarClose = (
    _event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSnackbarOpen(false);
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        <Typography textAlign={"center"} fontWeight={"700"} fontSize={"18px"}>
          Create account
        </Typography>
        <form onSubmit={handleSubmit(onSubmit)} className="Form">
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
          <TextField
            variant="outlined"
            placeholder="Password"
            type={showPassword ? "text" : "password"}
            {...register("password", {
              required: {
                value: true,
                message: "Please enter a password.",
              },
            })}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <LockOutlined />
                </InputAdornment>
              ),
              endAdornment: (
                <InputAdornment position="end">
                  {showPassword ? (
                    <VisibilityOff
                      onClick={() => setShowPassword(false)}
                      sx={{ cursor: "pointer" }}
                    />
                  ) : (
                    <Visibility
                      onClick={() => setShowPassword(true)}
                      sx={{ cursor: "pointer" }}
                    />
                  )}
                </InputAdornment>
              ),
            }}
          />
          {errors.password && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.password.message}
            </span>
          )}
          <Button
            variant="contained"
            color="primary"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Create account
          </Button>
        </form>
        <Typography textAlign={"center"}>
          Already have an account?{" "}
          <Link className="Link" href="/login">
            Sign in
          </Link>
        </Typography>
        <Divider sx={{ fontWeight: "700" }}>OR</Divider>
        {socialList.map((s) => (
          <Social key={s.name} type={s} />
        ))}
        <Divider />
        <Typography
          sx={{ textAlign: "center", fontSize: "12px", cursor: "pointer" }}
          onClick={() => setOpen(true)}
        >
          Personal Information Collection Statement
        </Typography>
      </Card>
      <Privacy open={open} handleClose={() => setOpen(false)} />
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "center",
        }}
      >
        <Alert
          onClose={handleSnackbarClose}
          severity="error"
          sx={{ width: "100%" }}
        >
          {error || "Something went wrong, unable to signup."}
        </Alert>
      </Snackbar>
    </div>
  );
}
