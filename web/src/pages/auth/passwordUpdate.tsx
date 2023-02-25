import { LockOutlined, Visibility, VisibilityOff } from "@mui/icons-material";
import {
  Alert,
  Button,
  Card,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import RentHomeLogo from "../../components/rentHomeLogo";
import { API_ADDRESS } from "../../config";
import { getTokenFromLocalStorage } from "../../helpers/auth";
import "../../styles/index.css";
import { ErrorResponse } from "../../types/types";

interface IPasswordUpdateInput {
  current_password: string;
  new_password: string;
}

export default function PasswordUpdatePage() {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<IPasswordUpdateInput>();
  const [showCurrentPassword, setShowCurrentPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const navigate = useNavigate();
  const [error, setError] = useState("");
  const [successful, setSuccessful] = useState(false);

  const onSubmit: SubmitHandler<IPasswordUpdateInput> = async (input) => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/change-password`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify(input),
      });
      if (res.ok) {
        setSuccessful(true);
      } else {
        const data: ErrorResponse = await res.json();
        setError(data.message);
      }
    } catch (error) {
      setError("Something went worng, unable to update password.");
    }
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        {successful ? (
          <div>
            <Typography
              textAlign={"center"}
              fontWeight={"700"}
              fontSize={"18px"}
            >
              Your password has been updated
            </Typography>
            <div style={{ padding: "20px 30px 20px 30px" }}>
              <Button
                variant="contained"
                sx={{ textTransform: "none", fontWeight: "700" }}
                fullWidth
                onClick={() => navigate("/me")}
              >
                Done
              </Button>
            </div>
          </div>
        ) : (
          <div>
            <Typography
              textAlign={"center"}
              fontWeight={"700"}
              fontSize={"18px"}
            >
              Update your password
            </Typography>
            <Typography
              textAlign={"center"}
              paddingTop="10px"
              paddingBottom="10px"
            >
              Your password must contain at least 8 characters
            </Typography>
            <form onSubmit={handleSubmit(onSubmit)} className="Form">
              {error && <Alert severity="warning">{error}</Alert>}
              <TextField
                variant="outlined"
                placeholder="Current password"
                type={showCurrentPassword ? "text" : "password"}
                {...register("current_password", {
                  required: {
                    value: true,
                    message: "Current password is required.",
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
                      {showCurrentPassword ? (
                        <VisibilityOff
                          onClick={() => setShowCurrentPassword(false)}
                          sx={{ cursor: "pointer" }}
                        />
                      ) : (
                        <Visibility
                          onClick={() => setShowCurrentPassword(true)}
                          sx={{ cursor: "pointer" }}
                        />
                      )}
                    </InputAdornment>
                  ),
                }}
              />
              {errors.current_password && (
                <span style={{ color: "red", fontSize: "14px" }}>
                  {errors.current_password.message}
                </span>
              )}
              <TextField
                variant="outlined"
                placeholder="New password"
                type={showNewPassword ? "text" : "password"}
                {...register("new_password", {
                  required: {
                    value: true,
                    message: "New password must contain at least 8 characters.",
                  },
                  minLength: {
                    value: 8,
                    message: "New password must contain at least 8 characters.",
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
              <Button
                variant="contained"
                size="large"
                sx={{ textTransform: "none", fontWeight: "600" }}
                type="submit"
              >
                Update password
              </Button>
            </form>
            <div style={{ paddingTop: "14px", paddingBottom: "14px" }}>
              <Button
                sx={{ textTransform: "none", fontWeight: "700" }}
                fullWidth
                onClick={() => navigate("/me")}
              >
                Cancel
              </Button>
            </div>
          </div>
        )}
      </Card>
    </div>
  );
}
