import {
  ArrowForwardIos,
  Close,
  EmailOutlined,
  LockOutlined,
  ToggleOffOutlined,
  ToggleOnOutlined,
} from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Dialog,
  Divider,
  IconButton,
  Snackbar,
  Typography,
} from "@mui/material";
import { red } from "@mui/material/colors";
import { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { API_ADDRESS } from "../config";
import { UserContext } from "../context/user";
import {
  getTokenFromLocalStorage,
  removeTokenFromLocalStorage,
  removeUserFromLocalStorage,
} from "../helpers/auth";
import { isNotificationState, isPrivacyState } from "../helpers/helper";
import { ErrorResponse, NotificationState, PrivacyState } from "../types/types";

export function MePage() {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();
  const [error, setError] = useState("");
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const { setUser } = useContext(UserContext);
  const [notifications, setNotifications] = useState<NotificationState[]>([]);
  const [privacies, setPrivacies] = useState<PrivacyState[]>([]);
  const [modalOpen, setModalOpen] = useState(false);
  const [item, setItem] = useState<NotificationState | PrivacyState>();

  useEffect(() => {
    // fetch user notifications
    (async function () {
      try {
        const res = await fetch(`${API_ADDRESS}/notifications`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ user_id: `${user?.id}` }),
        });

        if (res.ok) {
          const data = await res.json();
          setNotifications(data.notifications);
        } else {
          const data = await res.json();
          setError(data.message);
          setOpenSnackbar(true);
        }
      } catch (error) {
        setOpenSnackbar(true);
      }
    })();

    // fetch user privacies
    (async function () {
      try {
        const res = await fetch(`${API_ADDRESS}/privacies`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ user_id: `${user?.id}` }),
        });

        if (res.ok) {
          const data = await res.json();
          setPrivacies(data.privacies);
        } else {
          const data = await res.json();
          setError(data.message);
          setOpenSnackbar(true);
        }
      } catch (error) {
        setOpenSnackbar(true);
      }
    })();
  }, [user]);

  const handleModalClose = () => {
    setItem(undefined);
    setModalOpen(false);
  };

  const handleSnackbarClose = () => {
    setOpenSnackbar(false);
  };

  const handleLogout = async () => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/logout`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify({ user_id: `${user?.id}` }),
      });
      if (!res.ok) {
        const data = await res.json();
        setError(data.message);
        setOpenSnackbar(true);
      } else {
        removeUserFromLocalStorage();
        removeTokenFromLocalStorage();
        setUser(undefined);
        setOpenSnackbar(false);
      }
    } catch (error) {
      setOpenSnackbar(true);
    }
  };

  const updateNotifications = (notificationState: NotificationState) => {
    const newNotifications = notifications.map((ns) => {
      if (ns.notification.id === notificationState.notification.id) {
        ns.state = notificationState.state;
      }
      return ns;
    });
    setNotifications(newNotifications);
  };

  const updatePrivacies = (privacyState: PrivacyState) => {
    const newPrivacies = privacies.map((ps) => {
      if (ps.privacy.id === privacyState.privacy.id) {
        ps.state = privacyState.state;
      }
      return ps;
    });
    setPrivacies(newPrivacies);
  };

  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: "24px" }}>
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
          Account overview
        </Typography>
        <Button
          variant="contained"
          sx={{ textTransform: "none", fontWeight: "700" }}
          onClick={handleLogout}
        >
          Log out
        </Button>
      </div>
      <div>
        <Typography fontWeight={"600"} paddingBottom="10px">
          Email
        </Typography>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div>
            <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
              <EmailOutlined />
              <Typography>{"sudosanam@gmail.com"}</Typography>
            </div>
          </div>
          <Button
            sx={{ textTransform: "none", fontWeight: "bold" }}
            onClick={() => navigate("/email-update")}
          >
            Update
          </Button>
        </div>
        <Typography
          sx={{
            color: "rgb(92,158,141)",
            paddingLeft: "26px",
          }}
        >
          {"Verified"}
        </Typography>
      </div>
      <div>
        <Typography fontWeight={"600"} paddingBottom="10px">
          Password
        </Typography>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <LockOutlined />
            <Typography>{"********"}</Typography>
          </div>
          <Button
            sx={{ textTransform: "none", fontWeight: "bold" }}
            onClick={() => navigate("/password-update")}
          >
            Update
          </Button>
        </div>
      </div>
      <div>
        <Typography fontWeight={"600"}>{"Sign out on all devices"}</Typography>
        <Typography sx={{ paddingTop: "20px", paddingBottom: "20px" }}>
          {
            "Lost a device or recently used a public computer? Protect your account by signing out on all devices."
          }
        </Typography>
        <div>
          <Button
            variant="outlined"
            sx={{ textTransform: "none", fontWeight: "bold" }}
          >
            {" "}
            Sign out on all devices
          </Button>
        </div>
      </div>
      <div>
        <Button
          sx={{ color: red[500], textTransform: "none", fontWeight: "bold" }}
        >
          Delete account
        </Button>
      </div>
      <Divider />
      <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
        Notification settings
      </Typography>
      <Typography fontWeight={"600"} fontSize={"20px"}>
        Property journey
      </Typography>
      <div
        style={{
          borderRadius: "10px",
          border: "1px solid lightgray",
        }}
      >
        {notifications
          .filter((item) => item.notification.category === "Property journey")
          .map((item, i, arr) => (
            <div
              key={item.notification.id}
              style={{ cursor: "pointer" }}
              onClick={() => {
                setItem(item);
                setModalOpen(true);
              }}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  padding: "10px",
                }}
              >
                <div>
                  <Typography fontWeight={"600"}>
                    {item.notification.name}
                  </Typography>
                  <Typography>
                    {item.state + ": " + item.notification.method}
                  </Typography>
                </div>
                <ArrowForwardIos fontSize="small" />
              </div>
              {i < arr.length - 1 && <Divider />}
            </div>
          ))}
      </div>
      <Typography fontWeight={"600"} fontSize={"20px"}>
        Properties
      </Typography>
      <div
        style={{
          borderRadius: "10px",
          border: "1px solid lightgray",
        }}
      >
        {notifications
          .filter((item) => item.notification.category === "Properties")
          .map((item, i, arr) => (
            <div
              key={item.notification.id}
              style={{ cursor: "pointer" }}
              onClick={() => {
                setItem(item);
                setModalOpen(true);
              }}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  padding: "10px",
                }}
              >
                <div>
                  <Typography fontWeight={"600"}>
                    {item.notification.name}
                  </Typography>
                  <Typography>
                    {item.state + ": " + item.notification.method}
                  </Typography>
                </div>
                <ArrowForwardIos fontSize="small" />
              </div>
              {i < arr.length - 1 && <Divider />}
            </div>
          ))}
      </div>
      <Typography fontWeight={"600"} fontSize={"20px"}>
        Property market
      </Typography>
      <div
        style={{
          borderRadius: "10px",
          border: "1px solid lightgray",
        }}
      >
        {notifications
          .filter((item) => item.notification.category === "Property market")
          .map((item, i, arr) => (
            <div
              key={item.notification.id}
              style={{ cursor: "pointer" }}
              onClick={() => {
                setItem(item);
                setModalOpen(true);
              }}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  padding: "10px",
                }}
              >
                <div>
                  <Typography fontWeight={"600"}>
                    {item.notification.name}
                  </Typography>
                  <Typography>
                    {item.state + ": " + item.notification.method}
                  </Typography>
                </div>
                <ArrowForwardIos fontSize="small" />
              </div>
              {i < arr.length - 1 && <Divider />}
            </div>
          ))}
      </div>
      <Typography fontWeight={"600"} fontSize={"20px"}>
        Finance
      </Typography>
      <div
        style={{
          borderRadius: "10px",
          border: "1px solid lightgray",
        }}
      >
        {notifications
          .filter((item) => item.notification.category === "Finance")
          .map((item, i, arr) => (
            <div
              key={item.notification.id}
              style={{ cursor: "pointer" }}
              onClick={() => {
                setItem(item);
                setModalOpen(true);
              }}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  padding: "10px",
                }}
              >
                <div>
                  <Typography fontWeight={"600"}>
                    {item.notification.name}
                  </Typography>
                  <Typography>
                    {item.state + ": " + item.notification.method}
                  </Typography>
                </div>
                <ArrowForwardIos fontSize="small" />
              </div>
              {i < arr.length - 1 && <Divider />}
            </div>
          ))}
      </div>
      <Divider />
      <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
        Data privacy
      </Typography>
      <div
        style={{
          borderRadius: "10px",
          border: "1px solid lightgray",
        }}
      >
        {privacies.map((item, i) => {
          return (
            <div
              key={item.privacy.id}
              style={{ cursor: "pointer" }}
              onClick={() => {
                setItem(item);
                setModalOpen(true);
              }}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  padding: "10px",
                }}
              >
                <div>
                  <Typography fontWeight={"600"}>
                    {item.privacy.name}
                  </Typography>
                  <Typography>{item.state}</Typography>
                </div>
                <ArrowForwardIos fontSize="small" />
              </div>
              {i < privacies.length - 1 && <Divider />}
            </div>
          );
        })}
      </div>
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "center",
        }}
      >
        <Alert onClose={handleSnackbarClose} severity="error">
          {error || "Something went wrong, please try again."}
        </Alert>
      </Snackbar>
      {item && (
        <NotificationPrivacyUpdateModal
          item={item}
          open={modalOpen}
          onClose={handleModalClose}
          updateNotifications={updateNotifications}
          updatePrivacies={updatePrivacies}
        />
      )}
    </Box>
  );
}

const NotificationPrivacyUpdateModal = (props: {
  item: NotificationState | PrivacyState;
  open: boolean;
  onClose: () => void;
  updateNotifications: (notificationState: NotificationState) => void;
  updatePrivacies: (PrivacyState: PrivacyState) => void;
}) => {
  const { open, onClose, updateNotifications, updatePrivacies } = props;
  const [item, setItem] = useState(props.item);
  const [error, setError] = useState("");
  const initialState = props.item.state;

  const handleToggle = () => {
    setItem((prev) => {
      const newItem = { ...prev };
      if (prev.state === "On") {
        newItem.state = "Off";
      } else if (prev.state === "Off") {
        newItem.state = "On";
      }
      return newItem;
    });
  };

  const handleUpdate = async () => {
    setError("");
    try {
      if (isNotificationState(item)) {
        const res = await fetch(`${API_ADDRESS}/notifications/update`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${getTokenFromLocalStorage()}`,
          },
          body: JSON.stringify(item),
        });

        if (res.ok) {
          const data: NotificationState = await res.json();
          updateNotifications(data);
          onClose();
        } else {
          const data: ErrorResponse = await res.json();
          setError(data.message);
        }
      } else if (isPrivacyState(item)) {
        const res = await fetch(`${API_ADDRESS}/privacies/update`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${getTokenFromLocalStorage()}`,
          },
          body: JSON.stringify(item),
        });
        if (res.ok) {
          const data: PrivacyState = await res.json();
          updatePrivacies(data);
          onClose();
        } else {
          const data: ErrorResponse = await res.json();
          setError(data.message);
        }
      }
    } catch (error) {
      setError("Something went wrong, please try again.");
    }
  };

  return (
    <Dialog open={open} onClose={onClose}>
      {isNotificationState(item) && (
        <div style={{ padding: "32px 32px 0 32px" }}>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              gap: "20px",
            }}
          >
            <Typography sx={{ fontSize: "24px", fontWeight: "600" }}>
              {item.notification.name}
            </Typography>
            <IconButton onClick={onClose}>
              <Close fontSize="large" />
            </IconButton>
          </div>
          <Typography
            sx={{ paddingTop: "10px", paddingBottom: "10px", maxWidth: "22em" }}
          >
            {item.notification.description}
          </Typography>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Typography
              sx={{ display: "flex", alignItems: "center", gap: "8px" }}
            >
              <EmailOutlined color="primary" /> {item.notification.method}
            </Typography>
            <div
              style={{ paddingRight: "16px" }}
              onClick={() => {
                handleToggle();
                setError("");
              }}
            >
              {item.state === "On" && (
                <ToggleOnOutlined
                  sx={{ cursor: "pointer" }}
                  fontSize="large"
                  color="primary"
                />
              )}
              {item.state === "Off" && (
                <ToggleOffOutlined
                  sx={{ cursor: "pointer", opacity: "0.5" }}
                  fontSize="large"
                  color="primary"
                />
              )}
            </div>
          </div>
        </div>
      )}
      {isPrivacyState(item) && (
        <div style={{ padding: "32px 32px 10px 32px" }}>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Typography sx={{ fontSize: "24px", fontWeight: "600" }}>
              {item.privacy.name}
            </Typography>
            <IconButton onClick={onClose}>
              <Close fontSize="large" />
            </IconButton>
          </div>
          <Typography
            sx={{
              paddingTop: "10px",
              paddingBottom: "10px",
            }}
          >
            {item.privacy.description}
          </Typography>

          <div
            onClick={() => {
              handleToggle();
              setError("");
            }}
          >
            {item.state === "On" && (
              <ToggleOnOutlined
                sx={{ cursor: "pointer" }}
                fontSize="large"
                color="primary"
              />
            )}
            {item.state === "Off" && (
              <ToggleOffOutlined
                sx={{ cursor: "pointer", opacity: "0.5" }}
                fontSize="large"
                color="primary"
              />
            )}
          </div>
        </div>
      )}
      {error && (
        <Typography
          color={red[500]}
          textAlign="center"
          paddingTop={"10px"}
          paddingBottom={"10px"}
        >
          {error}
        </Typography>
      )}
      <Divider />
      <div style={{ padding: "20px 20px 20px 30px" }}>
        <Button
          variant="contained"
          disabled={initialState !== item.state ? false : true}
          onClick={handleUpdate}
          fullWidth
        >
          {"Save"}
        </Button>
      </div>
    </Dialog>
  );
};
