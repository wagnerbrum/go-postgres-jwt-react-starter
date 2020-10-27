import React, { useEffect } from "react";
import { Route, useHistory } from "react-router-dom";

import { isAuth } from "../../shared/utils";

const PrivateRoute = ({ component: Component, ...rest }) => {
    const history = useHistory();

    useEffect(() => {
        if (!isAuth()) {
            history.push("/");
        }
    }, []);

    return <Route {...rest} render={(props) => <Component {...props} />} />;
};

export default PrivateRoute;
