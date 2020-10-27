import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useForm } from "react-hook-form";

import { setToken } from "../../shared/utils";
import api from "../../services/api";

import "./index.css";

const Login = ({ history }) => {
    const [state, setState] = useState({
        isSubmitting: false,
        message: "",
    });
    const { register, handleSubmit, errors } = useForm();

    const { isSubmitting, message } = state;

    const onSubmit = async (data) => {
        setState({ ...state, isSubmitting: true, message: "" });

        const { input_email, input_password } = data;

        try {
            const res = await api.post("/login", {
                email: input_email,
                password: input_password,
            });

            const { token, user } = res.data;

            // expire in 30 minutes(same time as the cookie is invalidated on the backend)
            setToken(token);

            history.push({ pathname: "/home", state: user });
        } catch (e) {
            if (e.response && e.response.data) {
                const { msg } = e.response.data;

                return setState({
                    ...state,
                    message: msg,
                    isSubmitting: false,
                });
            }

            setState({ ...state, message: e.toString(), isSubmitting: false });
        }
    };

    return (
        <div className="container container-login">
            <div className="card text-center col-8 px-0">
                <div className="card-header">
                    <h2>Login</h2>
                </div>
                <div className="card-body">
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="row">
                            {message ? (
                                <div className="alert alert-warning col-8 mx-auto mt-2">{message}</div>
                            ) : (
                                <div className="alert alert-primary col-8 mx-auto mt-2">
                                    Enter with your credentials
                                </div>
                            )}
                        </div>

                        <div className="row mt-2">
                            <div className="form-group col-8 mx-auto text-left">
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Email"
                                    name="input_email"
                                    ref={register({ required: true })}
                                />
                                {errors.input_email && <small className="text-danger ">This field is required</small>}
                            </div>
                            <div className="form-group col-8 mx-auto text-left">
                                <input
                                    type="password"
                                    className="form-control"
                                    placeholder="Password"
                                    name="input_password"
                                    ref={register({ required: true })}
                                />
                                {errors.input_password && (
                                    <small className="text-danger ">This field is required</small>
                                )}
                            </div>
                        </div>

                        <button type="submit" className="btn btn-primary mx-auto my-2" disabled={isSubmitting}>
                            {isSubmitting ? "Loading ..." : "Login"}
                        </button>

                        <Link to="/register" className="mt-2">
                            Create your Account
                        </Link>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Login;
