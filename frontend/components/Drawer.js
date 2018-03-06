// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import MaterialDrawer from 'material-ui/Drawer';
import List, { ListItem, ListItemIcon, ListItemText } from 'material-ui/List';
import AddIcon from 'material-ui-icons/Add';
import StatisticsIcon from 'material-ui-icons/ShowChart';
import FitnessCenterIcon from 'material-ui-icons/FitnessCenter';
import Divider from 'material-ui/Divider';
import Link from 'next/link';

const styles = theme => ({
  list: {
    width: 250,
    background: theme.palette.background.paper,
  },
  listFull: {
    width: 'auto',
  },
});

type Props = {
  open: boolean,
  onRequestClose: Event => void,
  classes: Object,
};

function Drawer({ open, onRequestClose, classes }: Props) {
  const sideList = (
    <div className={classes.list}>
      <List>
        <Link prefetch href="/createMatch">
          <ListItem button>
            <ListItemIcon>
              <AddIcon />
            </ListItemIcon>
            <ListItemText primary="New Match" />
          </ListItem>
        </Link>
        <Divider />
        <Link prefetch href="/">
          <ListItem button>
            <ListItemIcon>
              <FitnessCenterIcon />
            </ListItemIcon>
            <ListItemText primary="Matches" />
          </ListItem>
        </Link>
        <Link prefetch href="/statistic">
          <ListItem button>
            <ListItemIcon>
              <StatisticsIcon />
            </ListItemIcon>
            <ListItemText primary="Statistics" />
          </ListItem>
        </Link>
      </List>
    </div>
  );

  return (
    <MaterialDrawer open={open} onClose={onRequestClose}>
      <div
        tabIndex={0}
        role="button"
        onClick={onRequestClose}
        onKeyDown={onRequestClose}
      >
        {sideList}
      </div>
    </MaterialDrawer>
  );
}

export default withStyles(styles)(Drawer);