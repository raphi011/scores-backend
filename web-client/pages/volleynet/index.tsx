import React from 'react';

import { QueryStringMapObject } from 'next';
import Router from 'next/router';

import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import Hidden from '@material-ui/core/Hidden';
import { createStyles, Theme, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import TournamentFilters, { Filters } from '../../components/volleynet/filters/TournamentFilters';
import TournamentList from '../../components/volleynet/TournamentList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { userSelector } from '../../redux/auth/selectors';
import {
  loadTournamentAction,
  loadTournamentsAction,
} from '../../redux/entities/actions';
import { filteredTournamentsSelector } from '../../redux/entities/selectors';
import { Store } from '../../redux/store';
import { Gender, Tournament, User } from '../../types';

const defaultLeagues = ['amateur-tour', 'pro-tour', 'junior-tour'];

const styles = (theme: Theme) => createStyles({
  primary: {
    flexGrow: 1,
  },
  root: {
    display: 'flex',

    flexDirection: 'row',

    [theme.breakpoints.down('xs')]: {
      flexDirection: 'column',
    }
  },
  secondary: {
    [theme.breakpoints.up('sm')]: {
      paddingRight: '30px',
      width: '300px',
    }
  },
});

interface Props extends WithStyles<typeof styles> {
  tournaments: Tournament[];
  league: string[];
  season: number;
  gender: Gender[];
  user: User;
  width: Breakpoint;

  loadTournaments: (
    filters: { gender: Gender[]; league: string[]; season: number },
  ) => void;
}

interface State {
  loading: boolean;
}

class Volleynet extends React.Component<Props, State> {
  static mapDispatchToProps = {
    loadTournament: loadTournamentAction,
    loadTournaments: loadTournamentsAction,
  };

  static buildActions({ gender, season, league = [] }: Props) {
      return [loadTournamentsAction({
        gender,
        league,
        season,
      })];
  }

  static getParameters(query: QueryStringMapObject) {
    const { season = "2018" } = query;
    let { gender = ['M', 'W'], league = ['amateur-tour'] } = query;

    if (!Array.isArray(league)) {
      league = [league];
    }
    if (!Array.isArray(gender)) {
      gender = [gender];
    }

    league = league.filter((l: string) => defaultLeagues.includes(l));

    return { league, gender, season: Number(season) };
  }

  static mapStateToProps(state: Store) {
    const tournaments = filteredTournamentsSelector(state);
    const { user } = userSelector(state);

    return { tournaments, user };
  }

  state = {
    loading: false,
  };


  renderList = (tournaments: Tournament[]) => {
    return (
      <div key={tournaments[0].id}>
        <TournamentList
          tournaments={tournaments}
          onTournamentClick={this.onTournamentClick}
        />
      </div>
    );
  };

  onTournamentClick = (t: Tournament) => {
    Router.push({
      pathname: '/volleynet/tournament',
      query: { id: t.id },
    });
  };

  onFilter = async (filters: Filters) => {
    this.setState({ loading: true })

    const { loadTournaments } = this.props;

    const query = filters; 

    Router.push({
      pathname: '/volleynet',
      query,
    });

    
    await loadTournaments(query);

    this.setState({ loading: false })
  }

  renderFilters = () => {
    const { league, gender, season, width } = this.props;
    const { loading } = this.state;

    const filters = (
      <TournamentFilters
        loading={loading}
        league={league}
        gender={gender}
        season={season}
        onFilter={this.onFilter}
      />
    )

    if (width === "xs") {
      return (
        <ExpansionPanel style={{ marginTop: '10px' }}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <Typography style={{ fontSize: '20px' }}>Filters</Typography>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails >
            {filters}
          </ExpansionPanelDetails>
        </ExpansionPanel>
      );
    }

    return filters;
  }

  render() {
    const { tournaments, classes } = this.props;

    return (
      <Layout title={{ text: 'Tournaments', href: '' }}>
        <div className={classes.root}>
          <div className={classes.secondary}>
            {this.renderFilters()}
          </div>
          <div className={classes.primary}>
            <GroupedList<Tournament>
              groupItems={groupTournaments}
              items={tournaments}
              renderHeader={renderHeader}
              renderList={this.renderList}
            />
          </div>
        </div>
      </Layout>
    );
  }
}

function sameDay(d1: Date, d2: Date): boolean {
  return (
    d1.getFullYear() === d2.getFullYear() &&
    d1.getMonth() === d2.getMonth() &&
    d1.getDay() === d2.getDay()
  );
}

function groupTournaments(tournaments: Tournament[]) {
  const grouped = [];

  let previous = null;

  tournaments.forEach(t => {
    if (!previous || !sameDay(new Date(previous.start), new Date(t.start))) {
      grouped.push([t]);
    } else {
      grouped[grouped.length - 1].push(t);
    }

    previous = t;
  });

  return grouped;
}

function renderHeader(tournaments: Tournament[]) {
  return (
    <DayHeader
      key={tournaments[0].start}
      appendix={`(${tournaments.length})`}
      date={new Date(tournaments[0].start)}
    />
  );
}

export default withStyles(styles)(withAuth(withWidth()(Volleynet)));
