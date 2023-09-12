import { LockOutlined, Visibility, VisibilityOff } from "@mui/icons-material";
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
import { useContext, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import Privacy from "../../components/privacy";
import RentHomeLogo from "../../components/rentHomeLogo";
import { API_ADDRESS } from "../../config";
import { UserContext } from "../../context/user";
import {
  getResetPasswordTokenFromLocalStorage,
  removeResetTokenFromLocalStorage,
  saveTokenInLocalStorage,
} from "../../helpers/auth";
import "../../styles/index.css";
import { ErrorResponse, User } from "../../types/types";

interface IConfirmForgotPasswordInput {
  code: string;
  new_password: string;
  confirm_password: string;
}

interface ConfirmForgotPasswordResponse {
  user: User;
  token: string;
}

export default function ConfirmForgotPasswordPage() {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<IConfirmForgotPasswordInput>();

  const [open, setOpen] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { setUser } = useContext(UserContext);

  const onSubmit: SubmitHandler<IConfirmForgotPasswordInput> = async (
    input
  ) => {
    try {
      const res = await fetch(
        `${API_ADDRESS}/api/auth/forgot-password-confirm`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${getResetPasswordTokenFromLocalStorage()}`,
          },
          body: JSON.stringify(input),
        }
      );
      if (res.ok) {
        const data: ConfirmForgotPasswordResponse = await res.json();
        saveTokenInLocalStorage(data.token);
        removeResetTokenFromLocalStorage();
        setUser(data.user);
        navigate("/");
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
          Update your password
        </Typography>
        <Typography textAlign={"center"}>
          We've sent a verification code to your email. Enter the code here,
          with your new password details.
        </Typography>
        <form className="Form" onSubmit={handleSubmit(onSubmit)}>
          {error && <Alert severity="warning">{error}</Alert>}
          <TextField
            variant="outlined"
            placeholder="Code"
            {...register("code", {
              required: {
                value: true,
                message: "Please enter a code.",
              },
            })}
          />
          {errors.code && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.code.message}
            </span>
          )}
          <TextField
            variant="outlined"
            placeholder="New password"
            type={showNewPassword ? "text" : "password"}
            {...register("new_password", {
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
                  {showNewPassword ? (
                    <VisibilityOff
                      onClick={() => setShowNewPassword(false)}
                      sx={{ cursor: "pointer" }}
                    />
                  ) : (
                    <Visibility
                      onClick={() => setShowNewPassword(true)}
                      sx={{ cursor: "pointer" }}
                    />
                  )}
                </InputAdornment>
              ),
            }}
          />
          {errors.new_password && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.new_password.message}
            </span>
          )}
          <TextField
            variant="outlined"
            placeholder="Confirm new password"
            type={showConfirmPassword ? "text" : "password"}
            {...register("confirm_password", {
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
                  {showConfirmPassword ? (
                    <VisibilityOff
                      onClick={() => setShowConfirmPassword(false)}
                      sx={{ cursor: "pointer" }}
                    />
                  ) : (
                    <Visibility
                      onClick={() => setShowConfirmPassword(true)}
                      sx={{ cursor: "pointer" }}
                    />
                  )}
                </InputAdornment>
              ),
            }}
          />
          {errors.confirm_password && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.confirm_password.message}
            </span>
          )}
          <Button
            variant="contained"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Change Password
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
