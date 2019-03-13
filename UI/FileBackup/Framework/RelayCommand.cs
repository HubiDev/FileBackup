using System;
using System.Windows.Input;

namespace FileBackup.Framework
{
    class RelayCommand<T> : ICommand
    {
        /// <summary>
        /// 
        /// </summary>
        public event EventHandler CanExecuteChanged;

        /// <summary>
        /// 
        /// </summary>
        private Action<T> _Action;

        /// <summary>
        /// 
        /// </summary>
        private Predicate<T> _CanExecutePredicate;

        /// <summary>
        /// 
        /// </summary>
        /// <param name="action"></param>
        /// <param name="canExecutePredicate"></param>
        public RelayCommand(Action<T> action, Predicate<T> canExecutePredicate = null)
        {
            _Action = action ?? throw new ArgumentNullException(action.ToString());
            _CanExecutePredicate = canExecutePredicate;
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="parameter"></param>
        /// <returns></returns>
        public bool CanExecute(object parameter)
        {
            return _CanExecutePredicate != null ? _CanExecutePredicate((T)parameter) : true;
        }

        public void Execute(object parameter)
        {
            _Action((T)parameter);
        }
    }
}
